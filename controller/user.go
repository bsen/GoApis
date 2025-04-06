package userController

type User struct {
	ID string 
	Name string
	Email string
	Age int64
	Phone string
}

type Users struct {
	Users []User
}

var usersJson = Users{
	Users: []User{
		{
		ID: "001",
		Name: "John Wick",
		Email: "john.wick@example.com",
		Age: 35,
		Phone: "123-456-7890",
		},
		{
		ID: "002",
		Name: "Peter Parker",
		Email: "peter.parker@example.com",
		Age: 28,
		Phone: "234-567-8901",
		},
		{
		ID: "003",
		Name: "John Wick",
		Email: "john.wick2@example.com",
		Age: 36,
		Phone: "345-678-9012",
		},
		{
		ID: "004",
		Name: "Peter Parker",
		Email: "peter.parker2@example.com",
		Age: 29,
		Phone: "456-789-0123",
		},
		{
		ID: "005",
		Name: "John Wick",
		Email: "john.wick3@example.com",
		Age: 37,
		Phone: "567-890-1234",
		},
		{
		ID: "006",
		Name: "Peter Parker",
		Email: "peter.parker3@example.com",
		Age: 30,
		Phone: "678-901-2345",
		},
		{
		ID: "007",
		Name: "John Wick",
		Email: "john.wick4@example.com",
		Age: 38,
		Phone: "789-012-3456",
		},
		{
		ID: "008",
		Name: "Peter Parker",
		Email: "peter.parker4@example.com",
		Age: 31,
		Phone: "890-123-4567",
		},
		{
		ID: "009",
		Name: "John Wick",
		Email: "john.wick5@example.com",
		Age: 39,
		Phone: "901-234-5678",
		},
		{
		ID: "010",
		Name: "Peter Parker",
		Email: "peter.parker5@example.com",
		Age: 32,
		Phone: "012-345-6789",
		},
		{
		ID: "011",
		Name: "Tony Stark",
		Email: "tony.stark@example.com",
		Age: 45,
		Phone: "111-222-3333",
		},
		{
		ID: "012",
		Name: "Bruce Wayne",
		Email: "bruce.wayne@example.com",
		Age: 40,
		Phone: "222-333-4444",
		},
		{
		ID: "013",
		Name: "Clark Kent",
		Email: "clark.kent@example.com",
		Age: 33,
		Phone: "333-444-5555",
		},
		{
		ID: "014",
		Name: "Diana Prince",
		Email: "diana.prince@example.com",
		Age: 3000,
		Phone: "444-555-6666",
		},
		{
		ID: "015",
		Name: "Steve Rogers",
		Email: "steve.rogers@example.com",
		Age: 104,
		Phone: "555-666-7777",
		},
		{
		ID: "016",
		Name: "Natasha Romanoff",
		Email: "natasha.romanoff@example.com",
		Age: 35,
		Phone: "666-777-8888",
		},
		{
		ID: "017",
		Name: "Stephen Strange",
		Email: "stephen.strange@example.com",
		Age: 42,
		Phone: "777-888-9999",
		},
		{
		ID: "018",
		Name: "Wade Wilson",
		Email: "wade.wilson@example.com",
		Age: 40,
		Phone: "888-999-0000",
		},
		{
		ID: "019",
		Name: "James Bond",
		Email: "james.bond@example.com",
		Age: 38,
		Phone: "999-000-1111",
		},
		{
		ID: "020",
		Name: "Sherlock Holmes",
		Email: "sherlock.holmes@example.com",
		Age: 40,
		Phone: "000-111-2222",
		},
	},
}

func GetUsersJson() Users{
	return usersJson
}