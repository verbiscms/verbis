import {HardBreak} from "tiptap-extensions";

export default class VerbisHardBreak extends HardBreak {

    get name() {
        return 'hard_break'
    }

    get schema() {
        return {
            inline: true,
            group: 'inline',
            selectable: false,
            parseDOM: [
                {tag: 'br'},
            ],
            toDOM: () => ['br'],
        }
    }
}
