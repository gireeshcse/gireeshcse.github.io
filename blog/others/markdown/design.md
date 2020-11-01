### Low Level Design

* extesible 

* Maintainability

    * What to change
    Self Explanatory
    Cohesion and coupling(SOLID)
    Strong cohesion and loose coupling

* Testable

Unit Testing
Integration Testing
A/B Testing
Stress Testing

How? 
OO Design
Abstraction
Design Patterns
    Sigleton/Buildup or Factory/Pub SUb/Chian of responsibility/adapter pattern
Proper Technologies
Testing

High level design 
* System architecture
*  scalability

Object Oriented

Class
    Attributes
    Behaviour
    Relationship
        Is(Pencil) - Inheritance,Interfaces
        Has(Battery,Graphite,Electronic)

Requiremnts Gathering
Use Cases
Actors/End Users
Scope - System Boudaries

Requirements
Work across cities
City can have multiple cinema halls
Search for currently active movies or past movies or upcomping
Filter by hall/city/genre


### Entities

* User

    - name
    - email
    - gender
    - dob(CBFC Rating of movies)
    - address
    - phone

    Feature
        - Book Ticket
        - Search movies
        - Give rating
        - see transaction history
        - Cancel/Resuduling
Movie

    - name
    - directors
    - description
    - audience_rating (13+)
    - length
    -actors
    - Language

    Behaviour

Booking
    Time
    No of seats
    total
    price
Coupons
    code
    expiry
    value
    max usage
Show
    timing
Hall
    no of seats 
    ac/non ac
    doors
    quality
Transaction
    date
    transaction_no
    payment_gateway
    amount
TrasactionStatus
    name
Profile
    name
    info 
    phot
    address
    phone
Rating
    value
    out of
    description
Genere
CIty
Seat
    - Physical Seat 
        recliner ,cup stand,cost, seat no
    - Show Seat - available or not , price
SeatSTatus
Cinema
    - Name
    - Address
    - No of halls
    - Has parking
    - Has Food facilities
Role
    name
    description
Admin
    name
    email
    could be a role

### Relationships

User
    User has a one Profile(1..1)
    user can have multiple Booking(1..*)
    User can have multiple Role(*..*)
Movie
     - 1..* Show
     - *..* Genre
Booking
    *..1 Show
    1..1 Rating
    1..* Transaction
    1..* Show Seat
    *..1 Coupon
Cinema 
    *-city
    1-* Hall