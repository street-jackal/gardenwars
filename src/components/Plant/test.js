const test = new Promise((resolve, reject) => {
  setTimeout(() => resolve("OK!"), 3000);
});

test.then((data) => console.log(data));
