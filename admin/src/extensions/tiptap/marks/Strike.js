import {Strike} from 'tiptap-extensions'

export default class VerbisStrike extends Strike {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            parseDOM: [
                {
                    tag: 's',
                },
                {
                    tag: 'del',
                },
                {
                    tag: 'strike',
                },
                {
                    style: 'text-decoration',
                    getAttrs: value => value === 'line-through',
                },
            ],
            toDOM: () => {
                if (this.options.tag) {
                    let attributes = this.options;
                    let tag = this.options.tag;
                    delete attributes.tag;
                    return attributes ? [tag, attributes, 0] : [tag, 0]
                } else {
                    return this.options ? ['s', this.options, 0] : ['s', 0]
                }
            },
        }
    }
}
