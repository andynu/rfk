#!/usr/bin/ruby env
# make a heatmap from a csv

require 'chunky_png'
require 'csv'

BLACK = ChunkyPNG::Color::BLACK
WHITE = ChunkyPNG::Color::WHITE

filename = ARGV[0]
height = `wc -l "#{filename}"`.strip.split(' ')[0].to_i + 1

image = nil
CSV.open(filename).each_with_index do |row, y|
  if image.nil?
    width = row.size+1
    image = ChunkyPNG::Image.new(width, height)
  end
  begin
    row.each_with_index do |val, x|
      if val.to_i > 0
        image[x,y] = BLACK
      else
        image[x,y] = WHITE
      end
    end
  rescue Exception => e
    warn e
  end
end

image.save("#{filename}.png")
