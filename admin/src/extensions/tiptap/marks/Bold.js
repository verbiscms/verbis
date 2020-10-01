import { Bold } from 'tiptap-extensions'

export default class VerbisBold extends Bold {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            parseDOM: [
                {
                    tag: 'strong',
                },
                {
                    tag: 'b',
                    getAttrs: node => node.style.fontWeight !== 'normal' && null,
                },
                {
                    style: 'font-weight',
                    getAttrs: value => /^(bold(er)?|[5-9]\d{2,})$/.test(value) && null,
                },
            ],
            toDOM: () => {

                if (this.options.tag) {
                    let attributes = this.options;
                    let tag = this.options.tag;
                    delete attributes.tag;
                    return attributes ? [tag, attributes, 0] : [tag, 0]
                } else {
                    return this.options ? ['strong', this.options, 0] : ['strong', 0]
                }
            },
        }
    }
}
