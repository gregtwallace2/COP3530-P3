<h1> Shakespeare’s Undying Love (of Words) </h1>

A straightforward program to examine word usage within Shakespeare’s works. The data set is sourced from
[here](https://bridgesdata.herokuapp.com/).

## Program Functions

When the program starts, it parses the json data set (~800,000 words) containing all of Shakespeare's works
and counts the use of each word. Punctuation is generally ignored, the apostrophes are assumed to be
part of words. Given some oddities in formatting in the data set, there is some room for cleanup here.

After the json data is parsed, individual words are loaded into a hash map (non-trivial struct 1) which
is used to effectively count the number of times each word is used. (The hash map maps words as keys to a
quantity of times the word is used in all of Shakespeare's works.)

After which, the main menu loads and offers two functions:

### Shakespeare's Top 50 Most Used Words

Function 1 produces a list of the top 50 most used words. The functionality is timed and offers a comparison
between loading the words from the hash map into a max heap (non-trivial struct 2) and loading the words into
a regular array and then using insertion sort to sort the values.

### Search Word for Usage Count

Function 2 allows the user to specify a word to search for and tells the user how many times Shakespeare
used that word. It provides a timing comparison between directly fetching the value from the hash map
and using the hash map to build a binary search tree and then fetching the answer from the binary search
tree.

#### Note on Timing

Of course, building additional data structures adds time. I tried to separate the timing of the build steps
from actual functions like searching or popping in order to more accurately have an apples to apples
comparison of times for similar functions.

Also, some functions are so fast, it appears they can't be accurately measured in go (results show 
0 nanoseconds). Additional work would need to be done if ultra high accuracy timing measurements were needed.

## Running the App

Install [Go](https://go.dev/dl/) version 1.25.4 or newer. Clone this repo to your local disk. From the
repo root, type `go run ./src`. Alternatively, the app can be compiled into an executable using the
command `go build ./src`.
