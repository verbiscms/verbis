/**
 * fieldParser.js
 * @author Ainsley Clark
 * @author URL:   https://reddico.co.uk
 * @author Email: ainsley@reddico.co.uk
 */

/**
 * FieldParser
 * Used for expanding and flattening fields
 * for the editor.
 *
 */
export default class FieldParser {
    /*
	 * constructor()
	 * Set fields and layout
	 */
    constructor(fields, layout) {
        this.fields = fields;
        this.layout = layout;
        this.parsed = {};
        this.flatFields = [];
    }
    /*
     * expandFields()
     * Expand the flat array of fields into an object for
     * the editor to process.
     */
    expandFields() {
        if (this.fields === null) {
            return {}
        }

        this.fields.forEach(field => {

            // Normal
            if (field.type !== "repeater" && !field.key.split("_").length) {
                this.parsed[field.name] = field;
                return
            }

            // Repeaters
            if (field.type === "repeater") {
                this._set(this.parsed, field.name + "_repeater", field)
                return
            }

            // Flexible
            if (field.type === "flexible") {
                this._set(this.parsed, field.name + "_flexible", field)
                return
            }

            // Obtain the split keys array
            const splitKeys = field.key.split("_");

            // Find repeater children
            const parent = this.fields.find(f => f.name === splitKeys[0]);
            if (parent && parent.type === "repeater") {
                for (let itemIndex = 1; itemIndex < splitKeys.length; itemIndex += 3) {
                    splitKeys.splice(itemIndex, 0, 'children');
                }
                this._set(this.parsed, splitKeys.join("_"), field)
            }

            // Find Flexible Children
            if (parent && parent.type === "flexible") {
                for (let itemIndex = 1; itemIndex < splitKeys.length; itemIndex += 3) {
                    splitKeys.splice(itemIndex, 0, 'children');
                }

                for (let itemIndex = 3; itemIndex < splitKeys.length; itemIndex += 3) {
                    splitKeys.splice(itemIndex, 0, 'fields');
                }

                parent.value.split(",").forEach((val, index) => {
                    let str = parent.name + "_children_" + index + "_type";
                    this._set(this.parsed, str, val)
                });

                this._set(this.parsed, splitKeys.join("_"), field)
            }
        });

        return this.parsed;
    }
    /*
     * flattenFields()
     * Collapse the fields into an array to send off
     * to the API.
     */
    flattenFields() {
        this._walker(this.fields);
        let fields = this.flatFields;
        this.flatFields = [];
        return fields;
    }
    /*
     * _walker()
     * Collapse the fields into an array to send off
     * to the API.
     */
    _walker(o) {
        if (Object.prototype.hasOwnProperty.call(o, "name")){
            this.flatFields.push(o);
        }
        for (const p in o) {
            if (Object.prototype.hasOwnProperty.call(o, p) && typeof o[p] === 'object' ) {
                if (o[p] !== null) {
                    this._walker(o[p]);
                }
            }
        }
    }
    /*
     * _set()
     * Vanilla version of Lodash's _set method,
     * "repeater_0_text" for example will turn into a nested object.
     */
    _set(obj, path, value){
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
