window.onload = () => {
    checkQuantityUpdatePrice()
    validateForms()
}

const searchProduct = async (customerId) => {
    const itemList = getItemList()

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(itemList),
    };
    const url = "/products?search=true&customerId=" + customerId

    await fetch(url, options)
    location.href = url;
}

const saveOrder = async (customerId) => {

    const finalPrice = document.querySelector("#finalPrice").textContent
    const orderId = document.getElementsByName("id")[0].value

    const itemList = getItemList()

    const customer = {
        id: customerId,
    }
    const order = {
        id: orderId,
        customer: customer,
        items: itemList,
        price: finalPrice,
        time: new Date().toISOString()
    }

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(order),
    };

    await fetch("/orders/save", options)
    location.href = "/orders"
}


const checkQuantityUpdatePrice = () =>{
    const list = getItemList()
    const finalPriceElem = document.querySelector("#finalPrice")
    let saveButtonDisabled = true;

    let saveOrderBtn = document.querySelector("#saveOrderBtn")

    let finalPrice = 0
    if (finalPriceElem != null){
        finalPriceElem.textContent = finalPrice.toString()
    }
    for (const i in list) {

        let span = document.querySelector("#error-" + list[i].index)

        finalPrice += list[i].price * list[i].quantity

        finalPriceElem.textContent = finalPrice.toString()

        if ((Number(list[i].quantity) > Number(list[i].stock)) ||
            (Number(list[i].quantity) <= 0 || list[i].quantity === "")) {
            span.classList.remove("d-none");
            span.classList.add("animate__headShake");
            finalPriceElem.textContent = "0";
            saveButtonDisabled = true
            break
        } else {
            if (span != null) {
                span.classList.add("d-none");
                span.classList.remove("animate__headShake");
            }
            saveButtonDisabled = false
        }
    }
    if (saveOrderBtn != null) {

        if (saveButtonDisabled) {
            saveOrderBtn.setAttribute("disabled", "")
        } else {
            saveOrderBtn.removeAttribute("disabled")

        }
    }

}

const validateDoc = () => {
    const docInput = document.getElementById("doc")
    if (docInput.value < 111111) {
        docInput.value = 111111
    }
}

const validatePhone = () => {
    const phoneInput = document.getElementById("phone")
    if (phoneInput.value < 100000) {
        phoneInput.value = 100000
    }
}

const validateQuantity = () => {
    const quantityInput = document.getElementById("quantity")
    if (quantityInput.value < 1) {
        quantityInput.value = 1
    }
}

const validatePrice = () => {
    const priceInput = document.getElementById("price")
    if (priceInput.value < 0.01) {
        priceInput.value = 1
    }
}

const validateForms = () => {
    const saveBtn = document.getElementById("saveBtn")
    const form = document.getElementById("form")

    let saveButtonDisabled = false;
    if ( form != null ) {
        const inputs = form.getElementsByTagName("input")
        for (let i in inputs) {
            if (i > 0) {
                console.log(inputs[i])
                console.log("i: " + i + " - " + inputs[i].value)
                if (inputs[i].value.length === 0) {
                    saveButtonDisabled = true
                    break
                }
                saveButtonDisabled = false
                console.log(inputs[i].value.length)
            }
        }
    }

    if (saveBtn != null) {
        if (saveButtonDisabled) {
            saveBtn.setAttribute("disabled", "")
        } else {
            saveBtn.removeAttribute("disabled")

        }
    }
}

const getItemList = () => {
    const itemsDivs = document.querySelectorAll(".item-info")
    const products = document.querySelectorAll(".card-subtitle")

    let list = []
    itemsDivs.forEach((div, key) => {
        let productQty = 0;
        products.forEach((prod, index) => {
            if (key === index) {
                const splitedStock = prod.textContent.split(": ")
                productQty = splitedStock[1].replace("\n", "").trim()
            }
        })
        const quantityInput = div.querySelector("input")
        const quantityValue = quantityInput.value
        const priceValue = div.querySelector(".d-none").textContent
        const productIdValue = quantityInput.getAttribute("id")

        const item = {
            index: key,
            productId: productIdValue,
            price: priceValue,
            quantity: quantityValue,
            stock: productQty
        }
        list.push(item)
    })
    return list

}

const eliminarItem = (element) => {

    let itemCard = element.closest(".card")

    itemCard.remove()
    checkQuantityUpdatePrice()
    refreshItems()
}

const refreshItems = () => {
    const itemsDivs = document.querySelectorAll(".item-info")
    itemsDivs.forEach((item, key) => {
        item.nextElementSibling.setAttribute("id", "error-" + key)
    })
}
