// _.isEmpty()

const MAX_SAFE_INTEGER = 9007199254740991;
const mapTag = '[object Map]',
    setTag = '[object Set]';
const objectProto = Object.prototype;
const hasOwnProperty = objectProto.hasOwnProperty;
const nativeObjectToString = objectProto.toString;
const symToStringTag = Symbol ? Symbol.toStringTag : undefined;
const nullTag = '[object Null]',
    undefinedTag = '[object Undefined]';
const asyncTag = '[object AsyncFunction]',
    funcTag = '[object Function]',
    genTag = '[object GeneratorFunction]',
    proxyTag = '[object Proxy]';
const isArray = Array.isArray;
const freeExports = typeof exports == 'object' && exports && !exports.nodeType && exports;
const freeModule = freeExports && typeof module == 'object' && module && !module.nodeType && module;
const moduleExports = freeModule && freeModule.exports === freeExports;
const freeGlobal = typeof global == 'object' && global && global.Object === Object && global;
const freeProcess = moduleExports && freeGlobal.process;
const freeSelf = typeof self == 'object' && self && self.Object === Object && self;
const root = freeGlobal || freeSelf || Function('return this')();
const Buffer = moduleExports ? root.Buffer : undefined;
const nativeIsBuffer = Buffer ? Buffer.isBuffer : undefined;
const isBuffer = nativeIsBuffer || stubFalse;
const propertyIsEnumerable = objectProto.propertyIsEnumerable;
const argsTag = '[object Arguments]',
    arrayTag = '[object Array]',
    boolTag = '[object Boolean]',
    dateTag = '[object Date]',
    errorTag = '[object Error]',
    numberTag = '[object Number]',
    objectTag = '[object Object]',
    regexpTag = '[object RegExp]',
    stringTag = '[object String]',
    weakMapTag = '[object WeakMap]';
const arrayBufferTag = '[object ArrayBuffer]',
    dataViewTag = '[object DataView]',
    float32Tag = '[object Float32Array]',
    float64Tag = '[object Float64Array]',
    int8Tag = '[object Int8Array]',
    int16Tag = '[object Int16Array]',
    int32Tag = '[object Int32Array]',
    uint8Tag = '[object Uint8Array]',
    uint8ClampedTag = '[object Uint8ClampedArray]',
    uint16Tag = '[object Uint16Array]',
    uint32Tag = '[object Uint32Array]';
const typedArrayTags = {};
typedArrayTags[float32Tag] = typedArrayTags[float64Tag] =
typedArrayTags[int8Tag] = typedArrayTags[int16Tag] =
typedArrayTags[int32Tag] = typedArrayTags[uint8Tag] =
typedArrayTags[uint8ClampedTag] = typedArrayTags[uint16Tag] =
typedArrayTags[uint32Tag] = true;
typedArrayTags[argsTag] = typedArrayTags[arrayTag] =
typedArrayTags[arrayBufferTag] = typedArrayTags[boolTag] =
typedArrayTags[dataViewTag] = typedArrayTags[dateTag] =
typedArrayTags[errorTag] = typedArrayTags[funcTag] =
typedArrayTags[mapTag] = typedArrayTags[numberTag] =
typedArrayTags[objectTag] = typedArrayTags[regexpTag] =
typedArrayTags[setTag] = typedArrayTags[stringTag] =
typedArrayTags[weakMapTag] = false;
const funcProto = Function.prototype;
const funcToString = funcProto.toString;
const promiseTag = '[object Promise]';
const dataViewCtorString = toSource(DataView),
    mapCtorString = toSource(Map),
    promiseCtorString = toSource(Promise),
    setCtorString = toSource(Set),
    weakMapCtorString = toSource(WeakMap);

function isObjectLike(value) {
  return value != null && typeof value == 'object';
}

function toSource(func) {
  if (func != null) {
    try {
      return funcToString.call(func);
    } catch (e) {}
    try {
      return (func + '');
    } catch (e) {}
  }
  return '';
}

const nodeUtil = (function() {
  try {
    // Use `util.types` for Node.js 10+.
    const types = freeModule && freeModule.require && freeModule.require('util').types;

    if (types) {
      return types;
    }

    // Legacy `process.binding('util')` for Node.js < 10.
    return freeProcess && freeProcess.binding && freeProcess.binding('util');
  } catch (e) {}
}());

const nodeIsTypedArray = nodeUtil && nodeUtil.isTypedArray;

function baseUnary(func) {
  return function(value) {
    return func(value);
  };
}

function baseIsTypedArray(value) {
  return isObjectLike(value) &&
    isLength(value.length) && !!typedArrayTags[baseGetTag(value)];
}

function baseIsArguments(value) {
  return isObjectLike(value) && baseGetTag(value) === argsTag;
}

const isArguments = baseIsArguments(function() { return arguments; }()) ? baseIsArguments : function(value) {
  return isObjectLike(value) && hasOwnProperty.call(value, 'callee') &&
    !propertyIsEnumerable.call(value, 'callee');
};

const isTypedArray = nodeIsTypedArray ? baseUnary(nodeIsTypedArray) : baseIsTypedArray;

function stubFalse() {
  return false;
}

function isLength(value) {
  return typeof value == 'number' &&
    value > -1 && value % 1 === 0 && value <= MAX_SAFE_INTEGER;
}

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

function isArrayLike(value) {
  return value != null && isLength(value.length) && !isFunction(value);
}

function isPrototype(value) {
  const Ctor = value && value.constructor,
      proto = (typeof Ctor == 'function' && Ctor.prototype) || objectProto;

  return value === proto;
}

function baseKeys(object) {
  if (!isPrototype(object)) {
    return nativeKeys(object);
  }
  const result = [];
  for (var key in Object(object)) {
    if (hasOwnProperty.call(object, key) && key !== 'constructor') {
      result.push(key);
    }
  }
  return result;
}

let getTag = baseGetTag;

if ((DataView && getTag(new DataView(new ArrayBuffer(1))) !== dataViewTag) ||
    (Map && getTag(new Map) !== mapTag) ||
    (Promise && getTag(Promise.resolve()) !== promiseTag) ||
    (Set && getTag(new Set) !== setTag) ||
    (WeakMap && getTag(new WeakMap) !== weakMapTag)) {
  getTag = function(value) {
    let result = baseGetTag(value),
        Ctor = result === objectTag ? value.constructor : undefined,
        ctorString = Ctor ? toSource(Ctor) : '';

    if (ctorString) {
      switch (ctorString) {
        case dataViewCtorString: return dataViewTag;
        case mapCtorString: return mapTag;
        case promiseCtorString: return promiseTag;
        case setCtorString: return setTag;
        case weakMapCtorString: return weakMapTag;
      }
    }
    return result;
  };
}

function isEmpty(value) {
  if (value == null) {
    return true;
  }
  if (isArrayLike(value) &&
      (isArray(value) || typeof value == 'string' || typeof value.splice == 'function' ||
        isBuffer(value) || isTypedArray(value) || isArguments(value))) {
    return !value.length;
  }
  const tag = getTag(value);
  if (tag === mapTag || tag === setTag) {
    return !value.size;
  }
  if (isPrototype(value)) {
    return !baseKeys(value).length;
  }
  for (let key in value) {
    if (hasOwnProperty.call(value, key)) {
      return false;
    }
  }
  return true;
}

console.log(isEmpty(null));  // => true
console.log(isEmpty(true));  // => true
console.log(isEmpty(1));  // => true
console.log(isEmpty([1, 2, 3]));  // => false
console.log(isEmpty({ 'a': 1 }));  // => false

/* Yandex решение */
function isLength(value) {
  return (
    typeof value === "number" &&
    value > -1 &&
    value % 1 === 0 &&
    value <= Number.MAX_SAFE_INTEGER
  );
}

function isNil(value) {
  return value === null || value === undefined;
}

function isArrayLike(value) {
  return !isNil(value) && typeof value !== "function" && isLength(value.length);
}

function isObjectLike(value) {
  return typeof value === "object" && value !== null;
}

function getTag(value) {
  if (value === null) {
    return value === undefined ? "[object Undefined]" : "[object Null]";
  }
  return toString.call(value);
}

const objectProto = Object.prototype;
function isPrototype(value) {
  const ctor = value && value.constructor;
  const proto = (typeof ctor === "function" && ctor.prototype) || objectProto;

  return value === proto;
}

function isArguments(value) {
  return isObjectLike(value) && getTag(value) === "[object Arguments]";
}

// Реализация лодаша
function isEmpty(value) {
  if (value === null) {
    return true;
  }

  if (
    isArrayLike(value) &&
    (Array.isArray(value) ||
      typeof value === "string" ||
      typeof value.splice === "function" ||
      isBuffer(value) ||
      isTypedArray(value) ||
      isArguments(value))
  ) {
    return !value.length;
  }

  const tag = getTag(value);
  if (tag === "[object Map]" || tag === "[object Set]") {
    return !value.size;
  }

  if (isPrototype(value)) {
    return !Object.keys(value).length;
  }

  for (const key in value) {
    if (hasOwnProperty.call(value, key)) {
      return false;
    }
  }

  return true;
}