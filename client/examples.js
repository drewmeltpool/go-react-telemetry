const { Client } = require("./tablets/client");

const client = Client("http://localhost:8080");

(async () => {
  let tablets;
  console.log("=== Scenario 1 ===");
  try {
    tablets = await client.ListTablets();
    console.log("Tablets:");
    console.table(tablets);
  } catch (err) {
    console.log(`Problem listing tablets: `, err);
  }
})();