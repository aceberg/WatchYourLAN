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
        someArray.sort(byFieldDown(field));
        oldField = field;
    } else {
        someArray.sort(byFieldUp(field));
        oldField = '';
    }

    if (field == 'State') {
        someArray.sort(byFieldUp('Watch'));
    }

    displayArrayData(someArray);
}

function byFieldUp(fieldName){
    return (a, b) => a[fieldName] < b[fieldName] ? 1 : -1;
}

function byFieldDown(fieldName){
    return (a, b) => a[fieldName] > b[fieldName] ? 1 : -1;
}