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
				console.log(this.options);
                return this.options ? ['a', {...node.attrs, ...this.options, target: '_blank'}, 0] : ['a', {...node.attrs, target: '_blank'}, 0];
            },
        }
    }
}
