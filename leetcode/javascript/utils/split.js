export function split(str, separator) {
    const res = [];
    let temp = '';

    for (let i = 0; i < str.length; i++) {
        const el = str[i];

        if (el === separator || separator === '' && temp) {
            res.push(temp);
            temp = '';
        }

        if (el !== separator) {
            temp += el;
        }
    }

    if (temp) {
        res.push(temp);
    }

    return res;
}