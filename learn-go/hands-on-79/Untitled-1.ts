// write a function to take in alpha numeric as string and return true or false for matching quotes {[()]}

function isAlphaNumeric(str) {
  return /^[a-z]+/.test(str);
}

const validPairs = {
  "{": "}",
  "[": "]",
  "(": ")",
};

function validateString(str: string): boolean {
  const stack: string[] = [];

  for (const s of str) {
    console.log(s);
    if (s in validPairs) {
      stack.push(s);
    } else if (Object.values(validPairs).includes(s)) {
      const value: string = stack.pop() || "";
      if (validPairs[value] !== s) {
        return false;
      }
    }
  }

  return stack.length === 0;
}

// test cases
console.log(validateString("{}[]()"));
console.log(validateString("{[}]"));
console.log(validateString("{([])}"));
console.log(validateString("{[])"));
console.log(validateString("((abc))"));
