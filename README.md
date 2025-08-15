# taxcal
A Tax Calculator on the commnad-line


# Assumptions & adjustment
- Edge cases like int overflowing have not been considered / addressed
- Slightly adjustment the output to be a table with more info


# Decisions
- Go 
    - I always find it very pleasant to read and trace other people's code in Go
    - it's compiled to a binary, easy and portable to test
- Tax Rates as a file
    - I'm aware that IRL this'll be from a DB and, from my experience in HSBC, this usually will be updated through a migration and pipeline under heavy review and regulation. But the main idea is decouples the data / configuration from the code
- Tests
    - For ease of implementation, I re-used the tax_rates.json as the test data, which I think is fine because real tax rates are not subject to change