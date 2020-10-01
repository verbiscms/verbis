import { CodeBlock } from "tiptap-extensions";

export default class VerbisCodeBlock extends CodeBlock {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            content: 'text*',
            marks: '',
            group: 'block',
            code: true,
            defining: true,
            draggable: false,
            parseDOM: [
                {tag: 'pre', preserveWhitespace: 'full'},
            ],
            toDOM: () => this.options ? ['pre', ['code', this.options, 0]] : ['pre', ['code', 0]],
        }
    }
}
