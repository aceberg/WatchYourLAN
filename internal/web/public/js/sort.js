var oldField = '';

function displayArrayData(someArray) {
    document.getElementById('tBody').innerHTML = "";

    let i = 0;
    for (let item of someArray){
        i = i + 1;
        html = createHTML(item, i);
        document.getElementById('tBody').insertAdjacentHTML('beforeend', html);
    }
}

function sortByAny(someArray, field) {

    if (field != oldField) {
        oldField = field;
        down = true;
    } else {
        oldField = '';
        down = false;
    }

    if (field == 'IP') {
        someArray.sort((a, b) => sortIP(a, b, down));
    } else {
        someArray.sort((a, b) => byField(a, b, field, down));
    }

    displayArrayData(someArray);
}

function byField(a, b, fieldName, down){
    if (a[fieldName] > b[fieldName]) {
        return down;
    } else {
        return !down;
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