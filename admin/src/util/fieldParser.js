/**
 * fieldParser.js
 * Custom Vue functions stored here.
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

/**
 * Require * Import
 *
 */

// eslint-disable-next-line no-unused-vars
export default class FieldParser {

    constructor(fields, layout) {
        this.fields = fields;
        this.layout = layout;
        this.parsed = {};
    }

    // Expand
    expandFields() {
        if (this.fields === null) {
            return {}
        }

        this.fields.forEach(field => {

            // Normal
            if (field.key === "" && field.type !== "repeater") {
                this.parsed[field.name] = field;
                return
            }

            this.set(this.parsed, field.key, field)
        });

        return this.parsed;
    }

    set(obj, path, value){
        if (Object(obj) !== obj) return obj; // When obj is not an object
        // If not yet an array, get the keys from the string-path
        if (!Array.isArray(path)) path = path.toString().match(/[^_[\]]+/g) || [];
        path.slice(0,-1).reduce((a, c, i) => // Iterate all of them except the last one
                Object(a[c]) === a[c] // Does the key exist and is its value an object?
                    // Yes: then follow that path
                    ? a[c]
                    // No: create the key. Is the next key a potential array-index?
                    : a[c] = Math.abs(path[i+1])>>0 === +path[i+1]
                    ? [] // Yes: assign a new array object
                    : {}, // No: assign a new plain object
            obj)[path[path.length-1]] = value; // Finally assign the value to the last key
        return obj; // Return the top-level object to allow chaining
    }
}
