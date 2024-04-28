/*                                                                          
|---------------------------------------------------------------            
| Inspire
|---------------------------------------------------------------            
|
| Generates a random inspirational quote 
| 
|
| @license: MIT
| @version: 1.0
| @since: 1.0
*/       
package main

import(
   "fmt"
   "math/rand"
   "time"
)

func main(){
   // Slice of quotes
	quotes := []string{
		"The only way to do great work is to love what you do. - Steve Jobs",
		"Believe you can and you're halfway there. - Theodore Roosevelt",
		"Success is not final, failure is not fatal: It is the courage to continue that counts. - Winston Churchill",
		"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
		"In the middle of difficulty lies opportunity. - Albert Einstein",
		"The only limit to our realization of tomorrow will be our doubts of today. - Franklin D. Roosevelt",
		"Do not wait for leaders; do it alone, person to person. - Mother Teresa",
		"You miss 100% of the shots you don't take. - Wayne Gretzky",
		"The greatest glory in living lies not in never falling, but in rising every time we fall. - Nelson Mandela",
		"You are never too old to set another goal or to dream a new dream. - C.S. Lewis",
		"The best time to plant a tree was 20 years ago. The second best time is now. - Chinese Proverb",
		"If you want to lift yourself up, lift up someone else. - Booker T. Washington",
		"The future depends on what you do today. - Mahatma Gandhi",
		"Don't watch the clock; do what it does. Keep going. - Sam Levenson",
		"Strive not to be a success, but rather to be of value. - Albert Einstein",
		"You must be the change you wish to see in the world. - Mahatma Gandhi",
		"Opportunities don't happen, you create them. - Chris Grosser",
		"Your time is limited, don't waste it living someone else's life. - Steve Jobs",
		"It does not matter how slowly you go as long as you do not stop. - Confucius",
		"The harder I work, the luckier I get. - Samuel Goldwyn",
		"Success is not the key to happiness. Happiness is the key to success. If you love what you are doing, you will be successful. - Albert Schweitzer",
		"The only person you should try to be better than is the person you were yesterday. - Anonymous",
		"A journey of a thousand miles begins with a single step. - Lao Tzu",
		"Success is walking from failure to failure with no loss of enthusiasm. - Winston Churchill",
		"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
		"Don't be pushed around by the fears in your mind. Be led by the dreams in your heart. - Roy T. Bennett",
		"The mind is everything. What you think you become. - Buddha",
		"I have not failed. I've just found 10,000 ways that won't work. - Thomas A. Edison",
		"In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln",
		"Don't let yesterday take up too much of today. - Will Rogers",
		"The only impossible journey is the one you never begin. - Tony Robbins",
		"The only way to do great work is to love what you do. - Steve Jobs",
		"You don't have to be great to start, but you have to start to be great. - Zig Ziglar",
		"The secret of getting ahead is getting started. - Mark Twain",
		"Life is 10% what happens to us and 90% how we react to it. - Charles R. Swindoll",
		"The purpose of our lives is to be happy. - Dalai Lama",
		"Dream big and dare to fail. - Norman Vaughan",
		"Don't count the days, make the days count. - Muhammad Ali",
		"The only thing standing between you and your goal is the story you keep telling yourself as to why you can't achieve it. - Jordan Belfort",
		"The best revenge is massive success. - Frank Sinatra",
		"Keep your face always toward the sunshine—and shadows will fall behind you. - Walt Whitman",
		"You are the average of the five people you spend the most time with. - Jim Rohn",
		"The man who has confidence in himself gains the confidence of others. - Hasidic Proverb",
		"You miss 100% of the shots you don't take. - Wayne Gretzky",
		"Failure is the condiment that gives success its flavor. - Truman Capote",
		"Believe you can and you're halfway there. - Theodore Roosevelt",
		"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
		"It always seems impossible until it's done. - Nelson Mandela",
		"Life is what happens when you're busy making other plans. - John Lennon",
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Call the getRandomQuote function and print the result
	fmt.Println(getRandomQuote(quotes))

}

func getRandomQuote(quotes []string) string {
	// Generate a random index within the range of the quotes slice
	randomIndex := rand.Intn(len(quotes))

	// Return the quote at the randomly generated index
	return quotes[randomIndex]
}
