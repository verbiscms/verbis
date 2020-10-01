import { Underline } from 'tiptap-extensions'

export default class VerbisUnderline extends Underline {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            parseDOM: [
                {
                    tag: 'u',
                },
                {
                    style: 'text-decoration',
                    getAttrs: value => value === 'underline',
                },
            ],
            toDOM: () => {
                if (this.options.tag) {
                    let attributes = this.options;
                    let tag = this.options.tag;
                    delete attributes.tag;
                    return attributes ? [tag, attributes, 0] : [tag, 0]
                } else {
                    return this.options ? ['u', this.options, 0] : ['u', 0]
                }
            },
        }
    }
}
