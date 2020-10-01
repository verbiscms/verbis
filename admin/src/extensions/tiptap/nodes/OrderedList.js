import { OrderedList } from "tiptap-extensions";

export default class VerbisOrderedList extends OrderedList {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            attrs: {
                order: {
                    default: 1,
                },
            },
            content: 'list_item+',
            group: 'block',
            parseDOM: [
                {
                    tag: 'ol',
                    getAttrs: dom => ({
                        order: dom.hasAttribute('start') ? +dom.getAttribute('start') : 1,
                    }),
                },
            ],
            toDOM: node => {
                if (this.options) {
                    return (node.attrs.order === 1 ? ['ol', this.options, 0] : ['ol', { start: node.attrs.order }, 0]);
                } else {
                    return (node.attrs.order === 1 ? ['ol', 0] : ['ol', { start: node.attrs.order }, 0]);
                }
            }
        }
    }

}
