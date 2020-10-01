import { Blockquote } from "tiptap-extensions";

export default class VerbisBlockquote extends Blockquote {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            content: 'block*',
            group: 'block',
            defining: true,
            draggable: false,
            parseDOM: [
                { tag: 'blockquote' },
            ],
            toDOM: () => this.options ? ['blockquote', this.options, 0] : ['blockquote', 0],
        }
    }
}
