process.nextTick(() => {
    console.log("1 tick");
    setTimeout(() => {
      console.log("in tick marco task")
    }, 0);
    // new Promise((res) => res())
      Promise.resolve()
        .then(() => console.log("in tick micro task1"))
        .then(() => console.log("in tick micro task2"));
      
    process.nextTick(()=>{
      console.log('in tick') 
    })
  });
  
  const bar = () => {
    console.log("bar");

  
      new Promise((res) => res())
  //   Promise.resolve()
      .then(() => console.log("anthor micro task1"))
      .then(() => console.log("anthor micro task2"));

      process.nextTick(() => {
        console.log("6 tick");
      });
  };
  
  const baz = () => console.log("baz");
  
  const foo = () => {
    console.log("foo");
  
    new Promise((resolve, reject) => {
      process.nextTick(() => {
        console.log("2 tick");
      });
      console.log(1);
      resolve("should be right after baz, before bar");
      console.log(2);
      return 3;
    })
      .then((resolve) => {
        process.nextTick(() => {
          console.log("4 tick");
        });
        console.log(resolve);
      })
      .then(() => {
        console.log("second job queue");
        process.nextTick(() => {
          console.log("5 tick");
        });
        setTimeout(() => {
          process.nextTick(() => {
            console.log("7 tick");
          });
          console.log("second message job");
        }, 0);
      });
    setTimeout(bar, 0);
    baz();
  };
  
  foo();
  
  process.nextTick(() => {
    console.log("3 tick");
  });

  