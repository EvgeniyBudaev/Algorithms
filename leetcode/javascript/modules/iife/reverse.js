window.invert = (function () {
    function reverse(str) {
        return str.split('').reverse().join('');
    }

    return {
        reverse,
    };
})();
