<?php
/**
 * MarginCrossOrderApiTest
 * PHP version 7.4
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Bitget Open API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 6.2.1
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Please update the test case below to test the endpoint.
 */

namespace Bitget\Test\Api;

use \Bitget\Configuration;
use \Bitget\ApiException;
use \Bitget\ObjectSerializer;
use Exception;
use PHPUnit\Framework\TestCase;

/**
 * MarginCrossOrderApiTest Class Doc Comment
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class MarginCrossOrderApiTest extends TestCase
{

    private $config;

    private $apiInstance;

    /**
     * Setup before running any test cases
     */
    public static function setUpBeforeClass(): void
    {
    }

    /**
     * Setup before running each test case
     */
    public function setUp(): void
    {
        $this->config = \Bitget\Config::getDefaultConfig();
        $this->apiInstance = new \Bitget\Api\MarginCrossOrderApi(
        // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
        // This is optional, `GuzzleHttp\Client` will be used as default.
            new \GuzzleHttp\Client(),
            $this->config
        );
    }

    /**
     * Clean up after running each test case
     */
    public function tearDown(): void
    {
    }

    /**
     * Clean up after running all test cases
     */
    public static function tearDownAfterClass(): void
    {
    }

    /**
     * Test case for marginCrossBatchCancelOrder
     *
     * batchCancelOrder.
     *
     */
    public function testMarginCrossBatchCancelOrder()
    {
        try {
            $req = new \Bitget\Model\MarginOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setSide("buy");
            $req->setOrderType("limit");
            $req->setTimeInForce("gtc");
            $req->setPrice("1600");
            $req->setBaseQuantity("0.625");
            $req->setQuoteAmount("1000");
            $req->setLoanType("normal");
            $result = $this->apiInstance->marginCrossPlaceOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderId());

            $req = new \Bitget\Model\MarginBatchCancelOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setOrderIds([$result->getData()->getOrderId()]);
            $result = $this->apiInstance->marginCrossBatchCancelOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getResultList()[0]->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossBatchFills
     *
     * fills.
     *
     */
    public function testMarginCrossBatchFills()
    {
        try {
            $result = $this->apiInstance->marginCrossFills("BTCUSDT","1679133422000");
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getFills());
            foreach ($result->getData()->getFills() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertNotNull($item->getSide());
                $this->assertNotNull($item->getOrderType());
                $this->assertNotNull($item->getOrderId());
                $this->assertNotNull($item->getFillTotalAmount());
                $this->assertNotNull($item->getFillQuantity());
                $this->assertNotNull($item->getFillPrice());
                $this->assertNotNull($item->getFeeCcy());
                $this->assertNotNull($item->getFees());
                $this->assertNotNull($item->getFillId());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossBatchHistoryOrders
     *
     * history.
     *
     */
    public function testMarginCrossBatchHistoryOrders()
    {
        try {
            $result = $this->apiInstance->marginCrossHistoryOrders("BTCUSDT","1679133422000");
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderList());
            foreach ($result->getData()->getOrderList() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertEquals($item->getSymbol(), "BTCUSDT");
                $this->assertNotNull($item->getSide());
                $this->assertNotNull($item->getSource());
                $this->assertNotNull($item->getStatus());
                $this->assertNotNull($item->getOrderId());
                $this->assertNotNull($item->getQuoteAmount());
                $this->assertNotNull($item->getBaseQuantity());
                $this->assertNotNull($item->getFillPrice());
                $this->assertNotNull($item->getFillQuantity());
                $this->assertNotNull($item->getFillTotalAmount());
                $this->assertNotNull($item->getLoanType());
                $this->assertNotNull($item->getOrderType());
                $this->assertNotNull($item->getPrice());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossBatchOpenOrders
     *
     * openOrders.
     *
     */
    public function testMarginCrossBatchOpenOrders()
    {
        try {
            $result = $this->apiInstance->marginCrossOpenOrders("BTCUSDT", "1679133422000" );
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderList());
            foreach ($result->getData()->getOrderList() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertEquals($item->getSymbol(), "BTCUSDT");
                $this->assertNotNull($item->getSide());
                $this->assertNotNull($item->getSource());
                $this->assertNotNull($item->getStatus());
                $this->assertNotNull($item->getOrderId());
                $this->assertNotNull($item->getQuoteAmount());
                $this->assertNotNull($item->getBaseQuantity());
                $this->assertNotNull($item->getFillPrice());
                $this->assertNotNull($item->getFillQuantity());
                $this->assertNotNull($item->getFillTotalAmount());
                $this->assertNotNull($item->getLoanType());
                $this->assertNotNull($item->getOrderType());
                $this->assertNotNull($item->getPrice());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossBatchPlaceOrder
     *
     * batchPlaceOrder.
     *
     */
    public function testMarginCrossBatchPlaceOrder()
    {
        try {
            $item = new \Bitget\Model\MarginOrderRequest(); //
            $item->setSymbol("BTCUSDT");
            $item->setSide("buy");
            $item->setOrderType("market");
            $item->setTimeInForce("gtc");
            $item->setQuoteAmount("10");
            $item->setLoanType("normal");
            $req = new \Bitget\Model\MarginBatchOrdersRequest();
            $req->setSymbol("BTCUSDT");
            $req->setOrderList([$item]);
            $result = $this->apiInstance->marginCrossBatchPlaceOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getResultList()[0]->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossCancelOrder
     *
     * cancelOrder.
     *
     */
    public function testMarginCrossCancelOrder()
    {
        try {
            $req = new \Bitget\Model\MarginOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setSide("buy");
            $req->setOrderType("limit");
            $req->setTimeInForce("gtc");
            $req->setPrice("1600");
            $req->setBaseQuantity("0.625");
            $req->setQuoteAmount("1000");
            $req->setLoanType("normal");
            $result = $this->apiInstance->marginCrossPlaceOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderId());

            $req = new \Bitget\Model\MarginCancelOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setOrderId($result->getData()->getOrderId());
            $result = $this->apiInstance->marginCrossCancelOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getResultList()[0]->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginCrossPlaceOrder
     *
     * placeOrder.
     *
     */
    public function testMarginCrossPlaceOrder()
    {
        try {
            $placeOrderReq = new \Bitget\Model\MarginOrderRequest(); //
            $placeOrderReq->setSymbol("BTCUSDT");
            $placeOrderReq->setSide("buy");
            $placeOrderReq->setOrderType("market");
            $placeOrderReq->setTimeInForce("gtc");
            $placeOrderReq->setQuoteAmount("10");
            $placeOrderReq->setLoanType("normal");
            $result = $this->apiInstance->marginCrossPlaceOrder($placeOrderReq);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }
}