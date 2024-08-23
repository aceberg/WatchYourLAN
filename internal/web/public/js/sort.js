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
    // console.log("Field =", field);

    if (field != oldField) {
        
        if (field == 'IP') {
            someArray.sort((a, b) => {
                const num1 = Number(a.IP.split(".").map((num) => (`000${num}`).slice(-3) ).join(""));
                const num2 = Number(b.IP.split(".").map((num) => (`000${num}`).slice(-3) ).join(""));
                return num1-num2;
            });
        } else {
            someArray.sort(byFieldDown(field));
        }

        oldField = field;
    } else {
        
        if (field == 'IP') {
            someArray.sort((a, b) => {
                const num1 = Number(a.IP.split(".").map((num) => (`000${num}`).slice(-3) ).join(""));
                const num2 = Number(b.IP.split(".").map((num) => (`000${num}`).slice(-3) ).join(""));
                return num2-num1;
            });
        } else {
            someArray.sort(byFieldUp(field));
        }

        oldField = '';
    }

    displayArrayData(someArray);
}

function byFieldUp(fieldName){
    return (a, b) => a[fieldName] < b[fieldName] ? 1 : -1;
}

function byFieldDown(fieldName){
    return (a, b) => a[fieldName] > b[fieldName] ? 1 : -1;
}