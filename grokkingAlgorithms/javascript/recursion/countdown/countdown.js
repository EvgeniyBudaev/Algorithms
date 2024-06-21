/**
 * Countdown
 * @param {number} i Number
 */
const countdown = i => {
    console.log(i);
    if (i <= 0) return 0;
    countdown(i-1);
};

countdown(3);