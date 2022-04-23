var storage = new Map();
const options = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
const baseURL = "http://tinyurl.com/";

var nextCode = () => {
    let key = "";
    for (let i = 0; i < 6; i++) {
        key += options.charAt(Math.floor(Math.random() * 62))
    }
    return key;
}
/**
 * Encodes a URL to a shortened URL.
 *
 * @param {string} longUrl
 * @return {string}
 */
var encode = function (longUrl) {
    let key = nextCode();
    while (storage.has(key)) {
        key = nextCode();
    }
    storage.set(key, longUrl);
    return `${baseURL}${key}`
};

/**
 * Decodes a shortened URL to its original URL.
 *
 * @param {string} shortUrl
 * @return {string}
 */
var decode = function (shortUrl) {
    const key = shortUrl.replace(baseURL, "");
    return storage.get(key)

};

/**
 * Your functions will be called as such:
 * decode(encode(url));
 */