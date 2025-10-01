# Exercise 2: Advanced GORM Operations
### Objective: Utilize GORM for more advanced operations including transactions, associations, and validation.
1.	Setup GORM with PostgreSQL:
      ○	Install GORM and the PostgreSQL driver.
      ○	Configure GORM with connection pooling.


2.	Create a Model with Associations:
      ○	Define a User model with fields and add an associated Profile model. For example:
      ■	User with fields: ID, Name, Age.
      ■	Profile with fields: ID, UserID, Bio, ProfilePictureURL.
      ○	Set up the one-to-one association between User and Profile.


3.	Auto Migrate with Constraints and Associations:
      ○	Use GORM’s AutoMigrate to create tables for User and Profile with appropriate constraints and associations.


4.	Insert Data with Associations:
      ○	Use GORM to insert a User and an associated Profile in a single transaction.


5.	Query Data with Associations:
      ○	Use GORM to retrieve users along with their profiles. Implement eager loading to optimize queries.


6.	Update and Delete Data:
      ○	Write functions to update a user’s profile and delete a user with associated profile, ensuring referential integrity.
