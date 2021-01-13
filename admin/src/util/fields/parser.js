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
            let keyArr = field.key.split("|");

            // Normal
            if (field.type !== "repeater" && field.type !== "flexible" && !field.key.includes("|")) {
                this.parsed[field.name] = field;
                return
            }

            if (field.type === "repeater") {
                if (field.key === "") {
                    this._set(this.parsed, field.name + "|repeater", field)
                    return;
                }

                return;
            }

            if (field.type === "flexible") {
                if (field.key === "") {
                    this._set(this.parsed, field.name + "|flexible", field)
                    return;
                }
            }

            keyArr.forEach((key, index) => {
                const f = this.fields.find(f => f.name === key);

                // TODO: Flexible Repeaters?

                if (f && field.type === "flexible") {
                    keyArr.splice(index + 2, 0, "flexible")
                }

                if (f && f.type === "repeater" && field.type !== "repeater") {
                    keyArr.splice(index + 1, 0, "children")
                }

                if (f && f.type === "flexible" && field.type !== "flexible") {
                    keyArr.splice(index + 1, 0, "children")
                }
            });

            this._set(this.parsed, keyArr.join("|"), field);
        });
        //
        //console.log(JSON.stringify(this.parsed, undefined, 2));

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
        if (!Array.isArray(path)) path = path.toString().match(/[^|[\]]+/g) || [];
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
