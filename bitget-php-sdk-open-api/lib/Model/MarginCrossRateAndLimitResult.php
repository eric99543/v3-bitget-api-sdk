<?php
/**
 * MarginCrossRateAndLimitResult
 *
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
 * Do not edit the class manually.
 */

namespace Bitget\Model;

use \ArrayAccess;
use \Bitget\ObjectSerializer;

/**
 * MarginCrossRateAndLimitResult Class Doc Comment
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class MarginCrossRateAndLimitResult implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'MarginCrossRateAndLimitResult';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'borrow_able' => 'bool',
        'coin' => 'string',
        'daily_interest_rate' => 'string',
        'leverage' => 'string',
        'max_borrowable_amount' => 'string',
        'transfer_in_able' => 'bool',
        'vips' => '\Bitget\Model\MarginCrossVipResult[]',
        'yearly_interest_rate' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'borrow_able' => null,
        'coin' => null,
        'daily_interest_rate' => null,
        'leverage' => null,
        'max_borrowable_amount' => null,
        'transfer_in_able' => null,
        'vips' => null,
        'yearly_interest_rate' => null
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'borrow_able' => false,
		'coin' => false,
		'daily_interest_rate' => false,
		'leverage' => false,
		'max_borrowable_amount' => false,
		'transfer_in_able' => false,
		'vips' => false,
		'yearly_interest_rate' => false
    ];

    /**
      * If a nullable field gets set to null, insert it here
      *
      * @var boolean[]
      */
    protected array $openAPINullablesSetToNull = [];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of nullable properties
     *
     * @return array
     */
    protected static function openAPINullables(): array
    {
        return self::$openAPINullables;
    }

    /**
     * Array of nullable field names deliberately set to null
     *
     * @return boolean[]
     */
    private function getOpenAPINullablesSetToNull(): array
    {
        return $this->openAPINullablesSetToNull;
    }

    /**
     * Setter - Array of nullable field names deliberately set to null
     *
     * @param boolean[] $openAPINullablesSetToNull
     */
    private function setOpenAPINullablesSetToNull(array $openAPINullablesSetToNull): void
    {
        $this->openAPINullablesSetToNull = $openAPINullablesSetToNull;
    }

    /**
     * Checks if a property is nullable
     *
     * @param string $property
     * @return bool
     */
    public static function isNullable(string $property): bool
    {
        return self::openAPINullables()[$property] ?? false;
    }

    /**
     * Checks if a nullable property is set to null.
     *
     * @param string $property
     * @return bool
     */
    public function isNullableSetToNull(string $property): bool
    {
        return in_array($property, $this->getOpenAPINullablesSetToNull(), true);
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'borrow_able' => 'borrowAble',
        'coin' => 'coin',
        'daily_interest_rate' => 'dailyInterestRate',
        'leverage' => 'leverage',
        'max_borrowable_amount' => 'maxBorrowableAmount',
        'transfer_in_able' => 'transferInAble',
        'vips' => 'vips',
        'yearly_interest_rate' => 'yearlyInterestRate'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'borrow_able' => 'setBorrowAble',
        'coin' => 'setCoin',
        'daily_interest_rate' => 'setDailyInterestRate',
        'leverage' => 'setLeverage',
        'max_borrowable_amount' => 'setMaxBorrowableAmount',
        'transfer_in_able' => 'setTransferInAble',
        'vips' => 'setVips',
        'yearly_interest_rate' => 'setYearlyInterestRate'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'borrow_able' => 'getBorrowAble',
        'coin' => 'getCoin',
        'daily_interest_rate' => 'getDailyInterestRate',
        'leverage' => 'getLeverage',
        'max_borrowable_amount' => 'getMaxBorrowableAmount',
        'transfer_in_able' => 'getTransferInAble',
        'vips' => 'getVips',
        'yearly_interest_rate' => 'getYearlyInterestRate'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->setIfExists('borrow_able', $data ?? [], null);
        $this->setIfExists('coin', $data ?? [], null);
        $this->setIfExists('daily_interest_rate', $data ?? [], null);
        $this->setIfExists('leverage', $data ?? [], null);
        $this->setIfExists('max_borrowable_amount', $data ?? [], null);
        $this->setIfExists('transfer_in_able', $data ?? [], null);
        $this->setIfExists('vips', $data ?? [], null);
        $this->setIfExists('yearly_interest_rate', $data ?? [], null);
    }

    /**
    * Sets $this->container[$variableName] to the given data or to the given default Value; if $variableName
    * is nullable and its value is set to null in the $fields array, then mark it as "set to null" in the
    * $this->openAPINullablesSetToNull array
    *
    * @param string $variableName
    * @param array  $fields
    * @param mixed  $defaultValue
    */
    private function setIfExists(string $variableName, array $fields, $defaultValue): void
    {
        if (self::isNullable($variableName) && array_key_exists($variableName, $fields) && is_null($fields[$variableName])) {
            $this->openAPINullablesSetToNull[] = $variableName;
        }

        $this->container[$variableName] = $fields[$variableName] ?? $defaultValue;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets borrow_able
     *
     * @return bool|null
     */
    public function getBorrowAble()
    {
        return $this->container['borrow_able'];
    }

    /**
     * Sets borrow_able
     *
     * @param bool|null $borrow_able borrow_able
     *
     * @return self
     */
    public function setBorrowAble($borrow_able)
    {

        if (is_null($borrow_able)) {
            throw new \InvalidArgumentException('non-nullable borrow_able cannot be null');
        }

        $this->container['borrow_able'] = $borrow_able;

        return $this;
    }

    /**
     * Gets coin
     *
     * @return string|null
     */
    public function getCoin()
    {
        return $this->container['coin'];
    }

    /**
     * Sets coin
     *
     * @param string|null $coin coin
     *
     * @return self
     */
    public function setCoin($coin)
    {

        if (is_null($coin)) {
            throw new \InvalidArgumentException('non-nullable coin cannot be null');
        }

        $this->container['coin'] = $coin;

        return $this;
    }

    /**
     * Gets daily_interest_rate
     *
     * @return string|null
     */
    public function getDailyInterestRate()
    {
        return $this->container['daily_interest_rate'];
    }

    /**
     * Sets daily_interest_rate
     *
     * @param string|null $daily_interest_rate daily_interest_rate
     *
     * @return self
     */
    public function setDailyInterestRate($daily_interest_rate)
    {

        if (is_null($daily_interest_rate)) {
            throw new \InvalidArgumentException('non-nullable daily_interest_rate cannot be null');
        }

        $this->container['daily_interest_rate'] = $daily_interest_rate;

        return $this;
    }

    /**
     * Gets leverage
     *
     * @return string|null
     */
    public function getLeverage()
    {
        return $this->container['leverage'];
    }

    /**
     * Sets leverage
     *
     * @param string|null $leverage leverage
     *
     * @return self
     */
    public function setLeverage($leverage)
    {

        if (is_null($leverage)) {
            throw new \InvalidArgumentException('non-nullable leverage cannot be null');
        }

        $this->container['leverage'] = $leverage;

        return $this;
    }

    /**
     * Gets max_borrowable_amount
     *
     * @return string|null
     */
    public function getMaxBorrowableAmount()
    {
        return $this->container['max_borrowable_amount'];
    }

    /**
     * Sets max_borrowable_amount
     *
     * @param string|null $max_borrowable_amount max_borrowable_amount
     *
     * @return self
     */
    public function setMaxBorrowableAmount($max_borrowable_amount)
    {

        if (is_null($max_borrowable_amount)) {
            throw new \InvalidArgumentException('non-nullable max_borrowable_amount cannot be null');
        }

        $this->container['max_borrowable_amount'] = $max_borrowable_amount;

        return $this;
    }

    /**
     * Gets transfer_in_able
     *
     * @return bool|null
     */
    public function getTransferInAble()
    {
        return $this->container['transfer_in_able'];
    }

    /**
     * Sets transfer_in_able
     *
     * @param bool|null $transfer_in_able transfer_in_able
     *
     * @return self
     */
    public function setTransferInAble($transfer_in_able)
    {

        if (is_null($transfer_in_able)) {
            throw new \InvalidArgumentException('non-nullable transfer_in_able cannot be null');
        }

        $this->container['transfer_in_able'] = $transfer_in_able;

        return $this;
    }

    /**
     * Gets vips
     *
     * @return \Bitget\Model\MarginCrossVipResult[]|null
     */
    public function getVips()
    {
        return $this->container['vips'];
    }

    /**
     * Sets vips
     *
     * @param \Bitget\Model\MarginCrossVipResult[]|null $vips vips
     *
     * @return self
     */
    public function setVips($vips)
    {

        if (is_null($vips)) {
            throw new \InvalidArgumentException('non-nullable vips cannot be null');
        }

        $this->container['vips'] = $vips;

        return $this;
    }

    /**
     * Gets yearly_interest_rate
     *
     * @return string|null
     */
    public function getYearlyInterestRate()
    {
        return $this->container['yearly_interest_rate'];
    }

    /**
     * Sets yearly_interest_rate
     *
     * @param string|null $yearly_interest_rate yearly_interest_rate
     *
     * @return self
     */
    public function setYearlyInterestRate($yearly_interest_rate)
    {

        if (is_null($yearly_interest_rate)) {
            throw new \InvalidArgumentException('non-nullable yearly_interest_rate cannot be null');
        }

        $this->container['yearly_interest_rate'] = $yearly_interest_rate;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset): bool
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    #[\ReturnTypeWillChange]
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value): void
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset): void
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    #[\ReturnTypeWillChange]
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


