<?php
/**
 * MarginPublicApiTest
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
use PHPUnit\Framework\TestCase;

/**
 * MarginPublicApiTest Class Doc Comment
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class MarginPublicApiTest extends TestCase
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
        $this->apiInstance = new \Bitget\Api\MarginPublicApi(
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
     * Test case for marginPublicCurrencies
     *
     * currencies.
     *
     */
    public function testMarginPublicCurrencies()
    {
        try {
            $result = $this->apiInstance->marginPublicCurrencies();
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            foreach ($result->getData() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertNotNull($item->getSymbol());
                $this->assertNotNull($item->getBaseCoin());
                $this->assertNotNull($item->getLiquidationRiskRatio());
                $this->assertNotNull($item->getMaxIsolatedLeverage());
                $this->assertNotNull($item->getPriceScale());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }
}
