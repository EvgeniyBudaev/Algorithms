/*
Генная строка может быть представлена 8-символьной длинной строкой с возможностью выбора из «A», «C», «G» и «T».

Предположим, нам нужно исследовать мутацию генной строки startGene в генную строку endGene, где одна мутация
определяется как изменение одного символа в генной строке.

Например, «AACCGGTT» --> «AACCGGTA» — это одна мутация.
Существует также банк генов, в котором регистрируются все действительные мутации генов. Чтобы генная строка была
действительной, ген должен находиться в банке.

Учитывая две строки генов startGene и endGene и банк генного банка, верните минимальное количество мутаций, необходимых
для мутации от startGene до endGene. Если такой мутации нет, верните -1.

Обратите внимание: предполагается, что отправная точка действительна, поэтому она может не быть включена в банк.
*/

/*
Input: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
Output: 1
 */

/*
Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
Output: 2
 */

/**
 * @param {string} startGene
 * @param {string} endGene
 * @param {string[]} bank
 * @return {number}
 */
var minMutation = function(startGene, endGene, bank) {
    const choices = ['A', 'C', 'G', 'T'];
    const queue = [[startGene, 0]];
    let seen = new Set();

    while (queue.length > 0) {
        const [gene, steps] = queue.shift();
        if (gene === endGene) return steps;
        for (let choice of choices) {
            for (let i = 0; i < gene.length; i++) {

                let mutatedGene = gene.slice(0, i) + choice + gene.slice(i + 1);
                if(bank.includes(mutatedGene) && !seen.has(mutatedGene)) {
                    seen.add(mutatedGene);
                    queue.push([mutatedGene, steps + 1]);
                }
            }
        }
    }

    return -1;
};

console.log(minMutation("AACCGGTT", "AACCGGTA", ["AACCGGTA"])); // 1
