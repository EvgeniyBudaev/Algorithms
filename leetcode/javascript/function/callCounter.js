const canGetCount = (number) => {
  let count = 0;

  return () => {
      count++;
    if (count <= number) {
      console.log("yes")
    } else {
      console.log("no")
    }
  };
};

const getOne = canGetCount(2);

getOne()  // yes
getOne()  // yes
getOne()  // no