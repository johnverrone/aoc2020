const fs = require('fs');

const nums = fs.readFileSync('input.txt').toString().split('\n').map(e => +e);

const challenge1 = () => {
  for (i in nums) {
    if (nums.includes(2020 - nums[i])) {
      return nums[i] * (2020 - nums[i]);
    }
  }
};

const challenge2 = () => {
  for (let i = 0; i < nums.length; i++) {
    const remainder = 2020 - nums[i];
    for (let j = i; j < nums.length; j++) {
      if (nums.includes(remainder - nums[j])) {
        return nums[i] * nums[j] * (remainder - nums[j]);
      }
    }
  }
}

console.log(challenge1());
console.log(challenge2());