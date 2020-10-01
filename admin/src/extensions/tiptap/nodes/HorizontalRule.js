import { HorizontalRule } from "tiptap-extensions";

export default class VerbisHorizontalRule extends HorizontalRule {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            group: 'block',
            parseDOM: [{tag: 'hr'}],
            toDOM: () => this.options ? ['hr', this.options] : ['hr'],
        }
    }

}
