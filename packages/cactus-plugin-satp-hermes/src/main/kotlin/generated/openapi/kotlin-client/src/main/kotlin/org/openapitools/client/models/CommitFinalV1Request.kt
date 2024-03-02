/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * 
 *
 * @param sessionID 
 * @param messageType 
 * @param clientIdentityPubkey 
 * @param serverIdentityPubkey 
 * @param commitFinalClaim 
 * @param hashCommitPrepareAck 
 * @param signature 
 * @param sequenceNumber 
 * @param commitFinalClaimFormat 
 * @param clientTransferNumber 
 */


data class CommitFinalV1Request (

    @Json(name = "sessionID")
    val sessionID: kotlin.String,

    @Json(name = "messageType")
    val messageType: kotlin.String,

    @Json(name = "clientIdentityPubkey")
    val clientIdentityPubkey: kotlin.String,

    @Json(name = "serverIdentityPubkey")
    val serverIdentityPubkey: kotlin.String,

    @Json(name = "commitFinalClaim")
    val commitFinalClaim: kotlin.String,

    @Json(name = "hashCommitPrepareAck")
    val hashCommitPrepareAck: kotlin.String,

    @Json(name = "signature")
    val signature: kotlin.String,

    @Json(name = "sequenceNumber")
    val sequenceNumber: java.math.BigDecimal,

    @Json(name = "commitFinalClaimFormat")
    val commitFinalClaimFormat: kotlin.Any? = null,

    @Json(name = "clientTransferNumber")
    val clientTransferNumber: kotlin.Int? = null

)
