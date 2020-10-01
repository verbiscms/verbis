import { Heading } from "tiptap-extensions";

export default class VerbisHeading extends Heading {

    constructor(levels, args) {
        super(levels);
        this.cmsOptions = args;
    }

    get schema() {
        return {
            attrs: {
                level: {
                    default: 1,
                },
            },
            content: 'inline*',
            group: 'block',
            defining: true,
            draggable: false,
            parseDOM: this.options.levels
                .map(level => ({
                    tag: `h${level}`,
                    attrs: {level},
                })),
            toDOM: node => {
                const attributes = this.cmsOptions[node.attrs.level];
                return attributes ? [`h${node.attrs.level}`, attributes, 0] : [`h${node.attrs.level}`, 0];
            }
        }
    }
}

