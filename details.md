Thank you for your interest in Gannett. The next step in our interview process is to complete a code sample to demonstrate your problem solving, coding and testing skills. There are basic requirements for the inputs and outputs but the design is completely open-ended and free for your implementation. There is no time limit on the assignment.

The solution should be delivered via a GitHub repository. The API application should be implemented in Golang. I recommend frequent pushing to GitHub. 

If you hit a technical hurdle in the exercise, please reach out to us to discuss.  There’s more to this exercise than just getting it right, it’s how you handle the journey.

You will be graded on quality of tests, code readability, and documentation for installation and execution.

Document any assumptions you make.

## Part 1: Supermarket API User Stories

You are a software developer employed by a local grocery store chain. The company needs an application which can add, delete, and fetch all produce in the system. 

| User Stories                | Narrative                                                                                 |
|-----------------------------|-------------------------------------------------------------------------------------------|
| Adding new produce          | As an employee, I want to add produce, so that I can add items to the database            |
| Deleting a produce item     | As an employee, I want to delete produce, so I can remove produce from the database       |
| Fetch the produce inventory | As an employee, I want to look up produce, so that I understand what produce is available |

### Acceptance Criteria
 * The produce database is only a single, in memory array of data and supports the following operations: `add`, `delete`, and `fetch`
 * Should support adding more than one new produce item at a time
 * The produce includes name, produce code, and unit price
 * The produce name is alphanumeric and case insensitive
 * The produce codes are sixteen characters long, with dashes separating each four character group
 * The produce codes are alphanumeric and case insensitive
 * The produce unit price is a number with up to 2 decimal places
 * Error handling (GET nonexistent produce, bad POST payload, etc.) 
 * RESTful API design, including utilizing proper response codes and correct HTTP Methods
 * Unit Testing
 * Concurrency - The main request should spawn off a go routine for writing / reading / deleting from the database. It will need to communicate data back to the main routine. An API request to add multiple items should do so concurrently.
 
 The API supports adding and deleting individual produce. You can also get any produce in the database.

#### GET
  Get all the produce in the database. Return JSON array of produce.


#### POST 
  Add one or more new produce items to the database. Accepts JSON input with the following parameters:
   * Name - required
   * Produce Code - required
   * Unit Price - required

   
#### DELETE 
  Delete a produce item from the database. Accepts a url parameter of `Produce Code` that will identify the item to delete.



Initially, the company has asked for the application to include the following produce in the new system:

| Produce Code        | Name         | Unit Price |
|---------------------|--------------|------------|
| A12T-4GH7-QPL9-3N4M | Lettuce      | $3.46      |
| E5T6-9UI3-TH15-QR88 | Peach        | $2.99      |
| YRT6-72AS-K736-L4AR | Green Pepper | $0.79      |
| TQ4C-VV6T-75ZX-1RMR | Gala Apple   | $3.59      |


## Part 2: Running the application

This application should be setup to run in a docker container. This will require a `Dockerfile` in your repository that your ci pipeline will use. When testing, we will pull your image from dockerhub and run it.

## (Optional) Part 3: Continuous Integration CI Pipeline

Use a cloud continuous integration platform such as GitHub Actions, TravisCI, CloudBee Jenkins, or CircleCI to create a continuous integration pipeline for your API.


 Upon a commit to Github automatically triggers the CI platform to:
 * build the code 
 * test the code
 * build the container
 * (Optional) push the container to dockerhub
 * (Optional) The CI server notifies the team of the build, test and/or deployment status via email, Slack, or some other means. 

Note: you don't have to actually deploy anywhere, but it would be worth thinking through what a deployment might look like 


## Conclusion 

We will walk through your code during your development (if necessary), along with when you have completed the assignment. We will be looking for you to show off both your technical and communication skills.
 
We would like you to complete the assignment as soon as you can.  Please keep us in the loop as to your progress. We want you to be comfortable and confident in the work you submit. Please let us know if you have any questions about this assignment or if you think of any questions about the position in advance of our review.