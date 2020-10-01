import { Italic } from 'tiptap-extensions'

export default class VerbisItalic extends Italic {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            parseDOM: [
                {tag: 'i'},
                {tag: 'em'},
                {style: 'font-style=italic'},
            ],
            toDOM: () => {
                if (this.options.tag) {
                    let attributes = this.options;
                    let tag = this.options.tag;
                    delete attributes.tag;
                    return attributes ? [tag, attributes, 0] : [tag, 0]
                } else {
                    return this.options ? ['em', this.options, 0] : ['em', 0]
                }
            },
        }
    }
}
