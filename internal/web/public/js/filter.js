let oldFilter = '';
let bkpArray;

function filterFunc(field, value) {

    if (oldFilter == field) {
        addrsArray = bkpArray;
    }

    switch (field) {
        case 'iface':
            addrsArray = addrsArray.filter((item) => item.Iface == value);
            break;
        case 'known':
            addrsArray = addrsArray.filter((item) => item.Known == value);
            break;
        case 'line':
            addrsArray = addrsArray.filter((item) => item.Now == value);
            break;
        default:
            console.log("Filter error");
    }
    
    oldFilter = field;

    displayArrayData(addrsArray);
}

function resetFilter() {
    addrsArray = bkpArray;
    displayArrayData(addrsArray);
}