class Customer
  @@no_of_customers=0
  def initialize(id,name,addr)
     @customer_id = id
     @customer_name = name
     @customer_addr = addr
  end
  def display_detail()
    puts "Customers id : #@customer_id"
    puts "Customers name : #@customer_name"
    puts "Customers address : #@customer_addr"
  end
  def total_no_of_customers()
    @@no_of_customers += 1
    puts "The numbers of all customers is : #@@no_of_customers"
  end
end

c1=Customer.new("1", "John", "Wisdom Apartments, Ludhiya")
c2=Customer.new("2", "Poul", "New Empire road, Khandala")
c1.display_detail()
c2.display_detail()