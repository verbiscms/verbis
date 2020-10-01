import { CodeBlockHighlight } from "tiptap-extensions";

export default class VerbisCodeBlockHighlight extends CodeBlockHighlight {

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
            toDOM: () => ['pre', ['code', 0]],
        }
    }
}
