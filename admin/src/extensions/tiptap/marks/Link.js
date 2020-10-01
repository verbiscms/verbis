import { Link } from 'tiptap-extensions'

export default class VerbisLink extends Link {

    constructor(args) {
        super();
        this.options = args;
    }

    get schema() {
        return {
            attrs: {
                href: {
                    default: null,
                },
            },
            inclusive: false,
            parseDOM: [
                {
                    tag: 'a[href]',
                    getAttrs: dom => ({
                        href: dom.getAttribute('href'),
                    }),
                },
            ],
            toDOM: node => {
                return this.options ? ['a', this.options, 0] : ['a', {...node.attrs, rel: 'noopener noreferrer nofollow'}, 0];
            },
        }
    }
}
