// _.rangeRight(4);

const MAX_SAFE_INTEGER = 9007199254740991;
const MAX_INTEGER = 1.7976931348623157e+308;
const INFINITY = 1 / 0;
const NAN = 0 / 0;
const reIsUint = /^(?:0|[1-9]\d*)$/;
const objectProto = Object.prototype;
const nativeObjectToString = objectProto.toString;
const nullTag = '[object Null]',
    undefinedTag = '[object Undefined]';
const symToStringTag = Symbol ? Symbol.toStringTag : undefined;
const asyncTag = '[object AsyncFunction]',
    funcTag = '[object Function]',
    genTag = '[object GeneratorFunction]',
    proxyTag = '[object Proxy]';
const symbolTag = '[object Symbol]';
const reIsBadHex = /^[-+]0x[0-9a-f]+$/i;
const reIsBinary = /^0b[01]+$/i;
const reIsOctal = /^0o[0-7]+$/i;
const reTrimStart = /^\s+/;
const reWhitespace = /\s/;
const freeParseInt = parseInt;
const nativeCeil = Math.ceil;
const nativeMax = Math.max;

function isObject(value) {
  const type = typeof value;
  return value != null && (type === 'object' || type === 'function');
}

function getRawTag(value) {
  const isOwn = hasOwnProperty.call(value, symToStringTag),
      tag = value[symToStringTag];

  try {
    value[symToStringTag] = undefined;
    var unmasked = true;
  } catch (e) {}

  const result = nativeObjectToString.call(value);
  if (unmasked) {
    if (isOwn) {
      value[symToStringTag] = tag;
    } else {
      delete value[symToStringTag];
    }
  }
  return result;
}

function objectToString(value) {
  return nativeObjectToString.call(value);
}

function baseGetTag(value) {
  if (value == null) {
    return value === undefined ? undefinedTag : nullTag;
  }
  return (symToStringTag && symToStringTag in Object(value))
    ? getRawTag(value)
    : objectToString(value);
}

function isFunction(value) {
  if (!isObject(value)) {
    return false;
  }
  // The use of `Object#toString` avoids issues with the `typeof` operator
  // in Safari 9 which returns 'object' for typed arrays and other constructors.
  const tag = baseGetTag(value);
  return tag === funcTag || tag === genTag || tag === asyncTag || tag === proxyTag;
}

function isLength(value) {
  return typeof value == 'number' &&
    value > -1 && value % 1 === 0 && value <= MAX_SAFE_INTEGER;
}

function isArrayLike(value) {
  return value != null && isLength(value.length) && !isFunction(value);
}

function isIndex(value, length) {
  const type = typeof value;
  length = length == null ? MAX_SAFE_INTEGER : length;

  return !!length &&
    (type === 'number' ||
      (type !== 'symbol' && reIsUint.test(value))) &&
        (value > -1 && value % 1 === 0 && value < length);
}

function eq(value, other) {
  return value === other || (value !== value && other !== other);
}

function isSymbol(value) {
  return typeof value == 'symbol' ||
    (isObjectLike(value) && baseGetTag(value) === symbolTag);
}

function trimmedEndIndex(string) {
  let index = string.length;

  while (index-- && reWhitespace.test(string.charAt(index))) {}
  return index;
}

function baseTrim(string) {
  return string
    ? string.slice(0, trimmedEndIndex(string) + 1).replace(reTrimStart, '')
    : string;
}

function toNumber(value) {
  if (typeof value == 'number') {
    return value;
  }
  if (isSymbol(value)) {
    return NAN;
  }
  if (isObject(value)) {
    let other = typeof value.valueOf == 'function' ? value.valueOf() : value;
    value = isObject(other) ? (other + '') : other;
  }
  if (typeof value != 'string') {
    return value === 0 ? value : +value;
  }
  value = baseTrim(value);
  const isBinary = reIsBinary.test(value);
  return (isBinary || reIsOctal.test(value))
    ? freeParseInt(value.slice(2), isBinary ? 2 : 8)
    : (reIsBadHex.test(value) ? NAN : +value);
}

function toFinite(value) {
  if (!value) {
    return value === 0 ? value : 0;
  }
  value = toNumber(value);
  if (value === INFINITY || value === -INFINITY) {
    let sign = (value < 0 ? -1 : 1);
    return sign * MAX_INTEGER;
  }
  return value === value ? value : 0;
}

function baseRange(start, end, step, fromRight) {
  let index = -1,
      length = nativeMax(nativeCeil((end - start) / (step || 1)), 0),
      result = Array(length);

  while (length--) {
    result[fromRight ? length : ++index] = start;
    start += step;
  }
  return result;
}

function isIterateeCall(value, index, object) {
  if (!isObject(object)) {
    return false;
  }
  const type = typeof index;
  if (type === 'number'
        ? (isArrayLike(object) && isIndex(index, object.length))
        : (type === 'string' && index in object)
      ) {
    return eq(object[index], value);
  }
  return false;
}

function createRange(fromRight) {
  return function(start, end, step) {
    if (step && typeof step != 'number' && isIterateeCall(start, end, step)) {
      end = step = undefined;
    }
    // Ensure the sign of `-0` is preserved.
    start = toFinite(start);
    if (end === undefined) {
      end = start;
      start = 0;
    } else {
      end = toFinite(end);
    }
    step = step === undefined ? (start < end ? 1 : -1) : toFinite(step);
    return baseRange(start, end, step, fromRight);
  };
}

const rangeRight = createRange(true);

console.log(rangeRight(4));  // => [3, 2, 1, 0]
console.log(rangeRight(-4));  // => [-3, -2, -1, 0]
console.log(rangeRight(1, 5));  // => [4, 3, 2, 1]
console.log(rangeRight(0, 20, 5));  // => [15, 10, 5, 0]
console.log(rangeRight(0, -4, -1));  // => [-3, -2, -1, 0]
console.log(rangeRight(1, 4, 0));  // => [1, 1, 1]
console.log(rangeRight(0));  // => []


/* Yandex решение */
const baseRange = (start, end, step, isRight) => {
  let index = -1;
  let length = Math.max(Math.ceil((end - start) / (step || 1)), 0);
  const result = new Array(length);

  while (length--) {
    result[isRight ? length : ++index] = start;
    start += step;
  }

  return result;
}

// Проверку на типы данных не добавлял, но студенты должны будут
function range(start = 0, end, step, isRight = false) {
        if (end === undefined) {
      end = start;
            start = 0;
    }

    step = step === undefined ? (start < end ? 1 : -1) : step;
    return baseRange(start, end, step, isRight);
}

function rangeRight(start, end, step) {
        return range(start, end, step, true);
}
