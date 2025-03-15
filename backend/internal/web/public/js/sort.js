let oldField = '';
let field = '';
let down = false;

function displayArrayData(someArray) {
    document.getElementById('tBody').innerHTML = "";

    let i = 0;
    for (let item of someArray){
        i = i + 1;
        html = createHTML(item, i);
        document.getElementById('tBody').insertAdjacentHTML('beforeend', html);
    }
}

function sortByAny(someArray) {

    if (field == 'IP') {
        someArray.sort((a, b) => sortIP(a, b, down));
    } else {
        someArray.sort((a, b) => byField(a, b, field, down));
    }

    displayArrayData(someArray);
}

function byField(a, b, fieldName, down){
    if (a[fieldName] > b[fieldName]) {
        return down ? 1 : -1;
    } else {
        return !down ? 1 : -1;
    }
}

function sortIP(a, b, down) {
    const num1 = numIP(a);
    const num2 = numIP(b);
    if (down) {
        return num1-num2;
    } else {
        return num2-num1;
    } 
}

function numIP(a) {
    return Number(a.IP.split(".").map((num) => (`000${num}`).slice(-3) ).join(""));
}