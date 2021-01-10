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

        this.layout.forEach(layout => {
            layout.fields.forEach(f => {
                this.resolveType(f)
            });
        });

        return this.parsed;
    }

    // Resolve
    resolveType(field) {
        switch (field.type) {

            // Repeaters
            case "repeater": {
                this.parsed[field.uuid] = this.getRepeater(field.uuid);
                break;
            }

            // Flexible
            case "flexible": {

                break;
            }

            // Default fields
            default: {
                const val = this.findByUUID(field.uuid)
                if (val) {
                   this.parsed[field.uuid] = val
                }
                break
            }
        }
    }

    findByUUID(uuid) {
        return this.fields.find(f => f.uuid === uuid)
    }

    getRepeater(uuid, index = null) {
        const repeater = this.fields.find(f => f.uuid === uuid && f.index === index);
        return {
            repeater: repeater,
            children: this.getRepeaterChildren(repeater.uuid, index)
        }
    }

    getRepeaterChildren(parent, index = null) {
        let arr = [];

        this.fields.filter(f => f.parent === parent).forEach(f => {
            arr[f.index] = arr[f.index] || {}
            if (f.type === "repeater") {

                const r = this.fields.find(t => t.uuid === f.uuid && t.index === f.index)
                arr[f.index][f.uuid] = {
                    repeater: r,
                    children: this.fields.filter(t => t.parent === f.uuid && r.index === f.index)
                }

                console.log(this.fields.filter(t => t.parent === f.uuid && r.index === f.index))
                return;
            }
            arr[f.index][f.uuid] = f;
        })
        console.log('IGNORE', index);
        return arr;
    }
}
