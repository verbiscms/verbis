import {Code} from 'tiptap-extensions'

export default class VerbisCode extends Code {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            excludes: '_',
            parseDOM: [
                {tag: 'code'},
            ],
            toDOM: () => {
                if (this.options.tag) {
                    let attributes = this.options;
                    let tag = this.options.tag;
                    delete attributes.tag;
                    return attributes ? [tag, attributes, 0] : [tag, 0]
                } else {
                    return this.options ? ['code', this.options, 0] : ['code', 0]
                }
            },
        }
    }
}
