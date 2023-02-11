search = (id) => {
    send("/search/" + id, null, (status, model) => {
        if (status === 200) {
            name.innerHTML = model.delivery.name
            phone.innerHTML = model.delivery.phone
            zip.innerHTML = model.delivery.zip
            address.innerHTML = model.delivery.address
            city.innerHTML = model.delivery.city
            region.innerHTML = model.delivery.region
            email.innerHTML = model.delivery.email

            order_id.innerHTML = model.order_uid
            track__number.innerHTML = model.track_number
            entry.innerHTML = model.entry

            locale.innerHTML = model.locale
            internal_signature.innerHTML = model.internal_signature
            customer_id.innerHTML = model.customer_id
            delivery_service.innerHTML = model.delivery_service
            shardkey.innerHTML = model.shardkey
            sm_id.innerHTML = model.sm_id
            date_created.innerHTML = model.date_created
            oof_shard.innerHTML = model.oof_shard
            locale.innerHTML = model.locale
            transaction.innerHTML = model.payment.transaction
            request_id.innerHTML = model.payment.request_id
            currency.innerHTML = model.payment.currency
            provider.innerHTML = model.payment.provider
            amount.innerHTML = model.payment.amount
            payment_dt.innerHTML = model.payment.payment_dt
            bank.innerHTML = model.payment.bank
            delivery_cost.innerHTML = model.payment.delivery_cost
            goods_total.innerHTML = model.payment.goods_total
            custom_fee.innerHTML = model.payment.custom_fee

            let items_
            if (items.children[0].id !== 'item0') {
                items_ = items.innerHTML
                items.innerHTML = ""
            } else {
                items_ = items.children[0].innerHTML
                items.innerHTML = ""
            }
            for (let i = 0; i < model.items.length; i++) {
                items.innerHTML += "<div id='item" + i + "'>" + items_ + "</div>"
                let item = items.querySelector("#item" + i)
                item.querySelector(".chrt_id").innerHTML = model.items[i].chrt_id
                item.querySelector(".track_number").innerHTML = model.items[i].track_number
                item.querySelector(".price").innerHTML = model.items[i].price
                item.querySelector(".rid").innerHTML = model.items[i].rid
                item.querySelector(".name").innerHTML = model.items[i].name
                item.querySelector(".sale").innerHTML = model.items[i].sale
                item.querySelector(".size").innerHTML = model.items[i].size
                item.querySelector(".total_price").innerHTML = model.items[i].total_price
                item.querySelector(".nm_id").innerHTML = model.items[i].nm_id
                item.querySelector(".brand").innerHTML = model.items[i].brand
                item.querySelector(".status").innerHTML = model.items[i].status
                item.querySelector(".order_uid").innerHTML = model.items[i].order_uid
            }
            order_bg.hidden = false
        } else {
            order_bg.hidden = true
            alert("Не найдено")
        }
    })
}

async function send(url, data, on_result) {
    let response;
    if (data !== undefined && data !== null) {
        response = await fetch(url, {
            method: "POST",
            body: data instanceof FormData ? data : JSON.stringify(data)
        });
    } else {
        response = await fetch(url);
    }
    let status = response.status;
    let message = await response.text();
    try {
        message = JSON.parse(message);
    } catch {
    }
    if (on_result) await on_result(status, message);
    return [status, message];
}