import { ListItem } from "tiptap-extensions";

export default class VerbisListItem extends ListItem {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            content: 'paragraph block*',
            defining: true,
            draggable: false,
            parseDOM: [
                { tag: 'li' },
            ],
            toDOM: () => this.options ? ['li', this.options, 0] : ['li', 0],
        }
    }
}
