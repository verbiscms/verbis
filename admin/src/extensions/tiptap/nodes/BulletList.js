import { BulletList } from "tiptap-extensions";

export default class VerbisBulletList extends BulletList {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            content: 'list_item+',
            group: 'block',
            parseDOM: [
                {
                    tag: 'ul'
                },
            ],
            toDOM: () => this.options ? ['ul', this.options, 0] : ['ul', 0],
        }
    }
}
