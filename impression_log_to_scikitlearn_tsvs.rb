#!/usr/bin/env ruby
require 'chronic'
require 'set'

# Convert impression log into linear set of tag impressions
#
# outputs env_details.tsv:
#    song  label  value  impression
#
env = {}
File.open('env_detail.tsv','w') do |f|
  File.open('data/impression.log').each_with_index do |line,i|
    time,tag,*rest = line.strip.split("\t")
    time = Chronic.parse(time)
    case tag
    when 'env'
      key, value = rest
      env[key] = value
    when 'tag'
      song, value, impression = rest
      f.puts [song, 'tag', value, impression].join("\t")
    when 'karma'
      song, impression = rest
      tags = {
        hour: time.hour,
        wday: time.wday,
        mday: time.mday,
      }
      tags.each_pair do |label,val|
        f.puts [song, label, val, impression].join("\t")
      end
      env.each_pair do |label,val|
        f.puts [song, label, val, impression].join("\t")
      end
    else
      warn [time,tag,rest].inspect
    end
  end
end



# aggregate details
hash = Hash.new{ |h,k| h[k] = Hash.new(&h.default_proc) }
File.open('env_detail.tsv').each_with_index do |line,i|
  song,label,value,impression = line.strip.split("\t")
  hash[song][label][value] = 0 if hash[song][label][value].kind_of? Hash
  hash[song][label][value] += 1
end

# output env_aggregate.tsv
#    song  label  value  sum_impression
File.open("env_aggregate.tsv",'w') do |f|
  hash.each_pair do |song,label_value_counts|
    label_value_counts.each_pair do |label,value_counts|
      value_counts.each_pair do |value,count|
        f.puts [song,label,value,count].join("\t")
      end
    end
  end
end


# Collect labels
labels = Set.new
song_label_impressions = Hash.new{ |h,k| h[k] = Hash.new(&h.default_proc) }
File.open('env_aggregate.tsv').each_with_index do |line,i|
  song,label,value,impression = line.strip.split("\t")
  key = [label,value].join(":")
  labels << key
  song_label_impressions[song][key] = impression
end

# Output group labels
labels = labels.to_a.sort
File.open("song_group_labels.tsv",'w') do |f|
  labels.each do |label|
    f.puts label
  end
end

# Output song impressions by label:
#   song  group1_impression  group2_impression  group3_impression  ...
File.open("song_group_matrix.tsv",'w') do |f|
  song_label_impressions.each_pair do |song, label_impressions|
    f.puts [song, labels.map{|key|
      (label_impressions[key] == {}) \
      ? 0 \
      : label_impressions[key]
    }].flatten.join("\t")
  end
end

