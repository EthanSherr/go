https://boot.dev/course/3b39d0f6-f944-4f1b-832d-a1daba32eda4/7faa7ccf-4fc0-406d-90b7-dac16e72f6c3/59b4509c-0cb1-487f-8b3f-68357629e60e

MULTIPLE INTERFACES
A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.

ASSIGNMENT
Complete the required methods so that the email type implements both the expense and printer interfaces.

COST()
If the email is not "subscribed", then the cost is 0.05 for each character in the body. If it is, then the cost is 0.01 per character.

Return the total cost of the entire email.

PRINT()
The print method should print to standard out the email's body text.