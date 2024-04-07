const express = require('express');
const cluster = require('node:cluster')

const app = express();
// The body-parser require is not needed for this example.
// const bodyparser = require("body-parser");
const PORT = process.env.PORT || 3000;

app.use(express.json()); // Corrected the missing parenthesis here
app.use(express.urlencoded({ extended: false }));

app.post('/api', (req, res) => {
  console.log(req.body);
  res.send(req.body);
});

app.get('/', (req, res) => {
  res.send("api chalne lagi");
});

app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});