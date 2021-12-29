  
  function test() {
    try {
      console.log(1);
      throw 3
      return 2;
    } catch (e){
      return e;
    } finally {
      return 4;
    }
  }
  
  console.log(test());