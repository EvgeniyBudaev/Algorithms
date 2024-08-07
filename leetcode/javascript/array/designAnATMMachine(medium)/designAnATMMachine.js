/* https://leetcode.com/problems/design-an-atm-machine/description/

Есть банкомат, в котором хранятся купюры 5 номиналов: 20, 50, 100, 200 и 500 долларов. Изначально банкомат пуст.
Пользователь может использовать машину для внесения или снятия любой суммы денег.

При снятии автомат отдает приоритет использованию банкнот большего номинала.

Например, если вы хотите снять 300 долларов и имеется 2 банкноты по 50 долларов, 1 банкнота по 100 долларов и 1 банкнота
по 200 долларов, то автомат будет использовать банкноты по 100 и 200 долларов.
Однако, если вы попытаетесь снять 600 долларов США и есть 3 банкноты по 200 долларов и 1 банкнота по 500 долларов, то
запрос на снятие будет отклонен, поскольку автомат сначала попытается использовать банкноту в 500 долларов, а затем не
сможет использовать банкноты для погашения оставшихся 100 долларов. Обратите внимание, что автомату не разрешено
использовать банкноты в 200 долларов вместо банкноты в 500 долларов.
Реализуйте класс ATM:

ATM() Инициализирует объект ATM.
void Deposit(int[] BanknotesCount) Вносит новые банкноты в порядке $20, $50, $100, $200 и $500.
int[]draw(int sum) Возвращает массив длиной 5, содержащий количество банкнот, которые будут переданы пользователю в
порядке $20, $50, $100, $200 и $500, и обновляет количество банкнот в банкомате после уход. Возвращает [-1], если это
невозможно (в этом случае не снимайте банкноты).

Input
["ATM", "deposit", "withdraw", "deposit", "withdraw", "withdraw"]
[[], [[0,0,1,2,1]], [600], [[0,1,0,1,1]], [600], [550]]
Output
[null, null, [0,0,1,0,1], null, [-1], [0,1,0,0,1]]

Объяснение
Банкомат = новый банкомат();
atm.deposit([0,0,1,2,1]); // Вносит 1 банкноту по 100 долларов, 2 банкноты по 200 долларов,
                          // и 1 банкнота номиналом 500 долларов.
банкомат.вывести(600);        // Возвращает [0,0,1,0,1]. В автомате используется 1 банкнота номиналом 100 долларов.
                          // и 1 банкнота номиналом 500 долларов. Банкноты, оставшиеся в
                          // машина — [0,0,0,2,0].
atm.deposit([0,1,0,1,1]); // Вносит 1 банкноту номиналом 50, 200 и 500 долларов.
                          // Банкноты в автомате теперь [0,1,0,3,1].
банкомат.вывести(600);        // Возвращает [-1]. Автомат попытается использовать банкноту в 500 долларов.
                          // и затем не сможем завершить оставшиеся 100 долларов,
                          // поэтому запрос на вывод средств будет отклонен.
                          // Так как запрос отклонен, количество банкнот
                          // в машине не модифицируется.
банкомат.вывести(550);        // Возвращает [0,1,0,0,1]. В автомате используется 1 банкнота номиналом 50 долларов.
                          // и 1 банкнота номиналом 500 долларов.
*/

var ATM = function() {
  this.banknotes = new Array(5).fill(0);
  this.billsList = {
    0: 20,
    1: 50,
    2: 100,
    3: 200,
    4: 500
  };
};

/**
 * @param {number[]} banknotesCount
 * @return {void}
 */
ATM.prototype.deposit = function(banknotesCount) {
  for (let i = 0; i <= 4; i++) {
    this.banknotes[i] += banknotesCount[i];
  }
};

/**
 * @param {number} amount
 * @return {number[]}
 */
ATM.prototype.withdraw = function(amount) {
  this.checkout = new Array(5).fill(0);

  for (let i = 4; i >= 0; i--) {
    let banknote = this.billsList[`${i}`];
    let division = Math.floor(amount / banknote);

    if (this.banknotes[i] - division < 0) {
      amount -= this.banknotes[i] * banknote
      this.checkout[i] = this.banknotes[i];
      this.banknotes[i] = 0;
    } else {
      amount -= banknote * division;
      this.banknotes[i] -= division;
      this.checkout[i] += division;
    }
  }
  if (amount !== 0) {
    this.deposit(this.checkout);
    return [-1];
  } else {
    return this.checkout;
  }
};

/**
 * Your ATM object will be instantiated and called as such:
 * var obj = new ATM()
 * obj.deposit(banknotesCount)
 * var param_2 = obj.withdraw(amount)
 */

const atm = new ATM();
atm.deposit([0,0,1,2,1]);
atm.withdraw(600);