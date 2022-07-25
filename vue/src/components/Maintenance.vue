<template>
    <div class="cover">
        <div class="popup-container">
            <div class="modal-body">
                Defund HQ Is Currently Undergoing Maintenance for the Testnet Stage 2 Upgrade
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: "Maintenance",
}
</script>

<style>
    .cover {
        position: fixed;
        top: 0;
        right: 0;
        bottom: 0;
        left: 0;
        z-index: 100;
        overflow: hidden;
        outline: 0;
        background-color: rgba(0, 0, 0, 0.185);
        display: flex;
        -webkit-box-pack: center;
        -webkit-justify-content: center;
        -ms-flex-pack: center;
        justify-content: center;
        -webkit-box-align: center;
        -webkit-align-items: center;
        -ms-flex-align: center;
        align-items: center;
    }
    .popup-container {
        font-size: medium;
        z-index: 1000;
        padding: 1.5rem;
        min-width: 300px;
        background-color: white;
        border: 0 solid rgba(0,0,0,.2);
        border-radius: 0.4375rem;
        max-width: 90%;
    }
    .close-div{
        position: absolute;
        width: 100vw;
        height: 100vh;
    }
    .modal-body {
        position: relative;
        flex: 1 1 auto;
        padding: 1.5rem;
    }
    .align-items-center {
        align-items: center !important;
    }
    .media {
        display: flex;
        align-items: flex-start;
    }
    .media-2 {
        display: flex;
        align-items: flex-start;
        margin-left: 0px;
    }
    .pic-size {
        width: 74px;
        height: 74px;
    }
    .pic {
        color: #fff;
        display: inline-flex;
        align-items: center;
        justify-content: center;
        font-size: 1rem;
        border-radius: 0.375rem;
        height: 75px;
        width: 75px;
    }
    .img {
        border-radius: 100px;
        height: -webkit-fill-available;
    }
    .details-div {
        margin-top: 15px;
        margin-bottom: 15px;
    }
    .desc-div {
        margin-top: 15px;
    }
    .details-within {
        font-weight: bold;
        margin-bottom: 10px;
    }
    .delegate-button-div {
        float: right;
        margin-top: 15px;
    }
    .list-group-item {
        position: relative;
        display: block;
        padding-top: 20px;
        padding-bottom: 20px;
        margin-bottom: -1px;
        background-color: #fff;
        border: 1px solid #e9ecef;
    }
    .list-group {
        display: flex;
        flex-direction: column;
        padding-left: 0;
        margin-bottom: 0;
    }
    .row {
        font-size: medium;
        border-color: rgba(0, 0, 0, 0.089);
        border-width: 2px;
        border-bottom-style: solid;
        padding-bottom: 25px;
    }
    .row-option {
        font-size: medium;
    }
    .voteOption {
        width: 15px;
        height: 15px;
        font-size: 15px;
        border-radius: 50%;
        border-width: 5px;
        border-style: solid;
        cursor: pointer;
        border-color: #b3b3b3;
    }
    .col-option {
        display: grid;
        grid-template-columns: 0.10fr 0.90fr;
    }
</style>