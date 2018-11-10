#!/usr/local/bin/ruby

$LOAD_PATH << "."
require "support"

class Decade
include Week
  @@num_of_year = 10
  def num_of_month
    puts Week::FIRST_DAY
    numbers = @@num_of_year * 12
    puts numbers
  end
end

d1 = Decade.new
d1.num_of_month()
Week.week_in_month()
Week.week_in_year()

